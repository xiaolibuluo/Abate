
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1298_0 {
    int value;
    struct Node1298_0 *next;
};

int length1298_0(const struct Node1298_0 *head) {
    int count = 0;
    const struct Node1298_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
#define MASK1298_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1298_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1298_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1298_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1298_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1298_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1298_3 { KIND_EMPTY1298_3 = 0, KIND_FULL1298_3 = 22 };

union Cell1298_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1298_3(union Cell1298_3 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL1298_3 : KIND_EMPTY1298_3;
}
#if defined(__STDC_VERSION__)
int variant1298_4(int value, ...) {
    return value + 29;
}
#else
int variant1298_4(int value) {
    return value + 29;
}
#endif
#define MASK1298_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1298_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1298_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
