
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2266_0(int value, ...) {
    return value + 72;
}
#else
int variant2266_0(int value) {
    return value + 72;
}
#endif
#if defined(__STDC_VERSION__)
int variant2266_1(int value, ...) {
    return value + 79;
}
#else
int variant2266_1(int value) {
    return value + 79;
}
#endif
#define MASK2266_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2266_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2266_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2266_3(int value, ...) {
    return value + 93;
}
#else
int variant2266_3(int value) {
    return value + 93;
}
#endif
struct Node2266_4 {
    int value;
    struct Node2266_4 *next;
};

int length2266_4(const struct Node2266_4 *head) {
    int count = 0;
    const struct Node2266_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
enum Kind2266_5 { KIND_EMPTY2266_5 = 0, KIND_FULL2266_5 = 10 };

union Cell2266_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2266_5(union Cell2266_5 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL2266_5 : KIND_EMPTY2266_5;
}
