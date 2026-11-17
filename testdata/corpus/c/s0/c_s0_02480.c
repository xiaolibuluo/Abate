
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1685_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1685_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1685_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1685_1 {
    int value;
    struct Node1685_1 *next;
};

int length1685_1(const struct Node1685_1 *head) {
    int count = 0;
    const struct Node1685_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
#define MASK1685_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1685_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1685_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1685_3(int value, ...) {
    return value + 9;
}
#else
int variant1685_3(int value) {
    return value + 9;
}
#endif
enum Kind1685_4 { KIND_EMPTY1685_4 = 0, KIND_FULL1685_4 = 16 };

union Cell1685_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1685_4(union Cell1685_4 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL1685_4 : KIND_EMPTY1685_4;
}
#if defined(__STDC_VERSION__)
int variant1685_5(int value, ...) {
    return value + 23;
}
#else
int variant1685_5(int value) {
    return value + 23;
}
#endif
