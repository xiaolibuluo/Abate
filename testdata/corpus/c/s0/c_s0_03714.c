
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2511_0(int value, ...) {
    return value + 56;
}
#else
int variant2511_0(int value) {
    return value + 56;
}
#endif
#define MASK2511_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2511_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2511_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2511_2(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 70;
}
#define MASK2511_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2511_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2511_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2511_4 {
    int value;
    struct Node2511_4 *next;
};

int length2511_4(const struct Node2511_4 *head) {
    int count = 0;
    const struct Node2511_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 84;
}
enum Kind2511_5 { KIND_EMPTY2511_5 = 0, KIND_FULL2511_5 = 91 };

union Cell2511_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2511_5(union Cell2511_5 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL2511_5 : KIND_EMPTY2511_5;
}
