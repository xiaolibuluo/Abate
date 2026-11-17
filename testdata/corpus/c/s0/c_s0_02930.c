
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1981_0 { KIND_EMPTY1981_0 = 0, KIND_FULL1981_0 = 53 };

union Cell1981_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1981_0(union Cell1981_0 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL1981_0 : KIND_EMPTY1981_0;
}
#if defined(__STDC_VERSION__)
int variant1981_1(int value, ...) {
    return value + 60;
}
#else
int variant1981_1(int value) {
    return value + 60;
}
#endif
struct Node1981_2 {
    int value;
    struct Node1981_2 *next;
};

int length1981_2(const struct Node1981_2 *head) {
    int count = 0;
    const struct Node1981_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
struct Node1981_3 {
    int value;
    struct Node1981_3 *next;
};

int length1981_3(const struct Node1981_3 *head) {
    int count = 0;
    const struct Node1981_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 74;
}
#define MASK1981_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1981_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1981_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1981_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1981_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1981_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
