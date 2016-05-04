---
layout: inner
title: ERROR invalid reference to FROM-clause
tags: ["postgresql"]
---
I was receiving this error:

**ERROR:  invalid reference to FROM-clause entry for table "mailboxer_conversations"
LINE 3: ... ON "mailboxer_notifications"."conversation_id" = "mailboxer...
                                                             ^
HINT:  There is an entry for table "mailboxer_conversations", but it cannot be referenced from this part of the query.**

Because I had a query like this:

{% highlight sql %}

SELECT  DISTINCT "mailboxer_conversations".* FROM "mailboxer_conversations",
(select count(*) from mailboxer_receipts where mailboxer_receipts.is_read = 't') as read_count
INNER JOIN "mailboxer_notifications" ON "mailboxer_notifications"."conversation_id" = "mailboxer_conversations"."id"
INNER JOIN "mailboxer_receipts" ON "mailboxer_receipts"."notification_id" = "mailboxer_notifications"."id"
WHERE "mailboxer_receipts"."receiver_id" = 118
ORDER BY "mailboxer_conversations"."id" ASC LIMIT 1000

{% endhighlight %}

Fixed it by rewriting the query like this:

{% highlight sql %}

SELECT  DISTINCT "mailboxer_conversations".*,
(select count(*) from mailboxer_receipts where mailboxer_receipts.is_read = 't') as read_count
FROM "mailboxer_conversations"
INNER JOIN "mailboxer_notifications" ON "mailboxer_notifications"."conversation_id" = "mailboxer_conversations"."id"
INNER JOIN "mailboxer_receipts" ON "mailboxer_receipts"."notification_id" = "mailboxer_notifications"."id"
WHERE "mailboxer_receipts"."receiver_id" = 118
ORDER BY "mailboxer_conversations"."id" ASC LIMIT 1000

{% endhighlight %}
