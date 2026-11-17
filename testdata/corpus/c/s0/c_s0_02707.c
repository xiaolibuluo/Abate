
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1840_0 { KIND_EMPTY1840_0 = 0, KIND_FULL1840_0 = 63 };

union Cell1840_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1840_0(union Cell1840_0 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL1840_0 : KIND_EMPTY1840_0;
}
#define MASK1840_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1840_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1840_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1840_2 { KIND_EMPTY1840_2 = 0, KIND_FULL1840_2 = 77 };

union Cell1840_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1840_2(union Cell1840_2 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL1840_2 : KIND_EMPTY1840_2;
}
struct Node1840_3 {
    int value;
    struct Node1840_3 *next;
};

int length1840_3(const struct Node1840_3 *head) {
    int count = 0;
    const struct Node1840_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 84;
}
#define MASK1840_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1840_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1840_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1840_5(int value, ...) {
    return value + 98;
}
#else
int variant1840_5(int value) {
    return value + 98;
}
#endif
enum Kind1840_6 { KIND_EMPTY1840_6 = 0, KIND_FULL1840_6 = 8 };

union Cell1840_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1840_6(union Cell1840_6 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL1840_6 : KIND_EMPTY1840_6;
}
