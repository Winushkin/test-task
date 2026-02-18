SELECT
    num, mqtt, inv_id, unit_guid,
    message_id, message_text, context, message_class,
    message_level, area, var_addr, block_sign,
    message_type, bit_number, invert_bit, file_id
FROM records
LIMIT @limit OFFSET @offset;