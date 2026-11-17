
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1637_0(int value, ...) {
    return value + 43;
}
#else
int variant1637_0(int value) {
    return value + 43;
}
#endif
enum Kind1637_1 { KIND_EMPTY1637_1 = 0, KIND_FULL1637_1 = 50 };

union Cell1637_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1637_1(union Cell1637_1 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL1637_1 : KIND_EMPTY1637_1;
}
#if defined(__STDC_VERSION__)
int variant1637_2(int value, ...) {
    return value + 57;
}
#else
int variant1637_2(int value) {
    return value + 57;
}
#endif
#define MASK1637_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1637_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1637_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1637_4 {
    int value;
    struct Node1637_4 *next;
};

int length1637_4(const struct Node1637_4 *head) {
    int count = 0;
    const struct Node1637_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}
#if defined(__STDC_VERSION__)
int variant1637_5(int value, ...) {
    return value + 78;
}
#else
int variant1637_5(int value) {
    return value + 78;
}
#endif
#define MASK1637_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1637_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1637_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
