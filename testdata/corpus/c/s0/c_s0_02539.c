
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1727_0 { KIND_EMPTY1727_0 = 0, KIND_FULL1727_0 = 49 };

union Cell1727_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1727_0(union Cell1727_0 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL1727_0 : KIND_EMPTY1727_0;
}
#define MASK1727_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1727_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1727_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1727_2(int value, ...) {
    return value + 63;
}
#else
int variant1727_2(int value) {
    return value + 63;
}
#endif
#if defined(__STDC_VERSION__)
int variant1727_3(int value, ...) {
    return value + 70;
}
#else
int variant1727_3(int value) {
    return value + 70;
}
#endif
struct Node1727_4 {
    int value;
    struct Node1727_4 *next;
};

int length1727_4(const struct Node1727_4 *head) {
    int count = 0;
    const struct Node1727_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 77;
}
