
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1105_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1105_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1105_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1105_1 {
    int value;
    struct Node1105_1 *next;
};

int length1105_1(const struct Node1105_1 *head) {
    int count = 0;
    const struct Node1105_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
enum Kind1105_2 { KIND_EMPTY1105_2 = 0, KIND_FULL1105_2 = 28 };

union Cell1105_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1105_2(union Cell1105_2 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1105_2 : KIND_EMPTY1105_2;
}
#if defined(__STDC_VERSION__)
int variant1105_3(int value, ...) {
    return value + 35;
}
#else
int variant1105_3(int value) {
    return value + 35;
}
#endif
enum Kind1105_4 { KIND_EMPTY1105_4 = 0, KIND_FULL1105_4 = 42 };

union Cell1105_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1105_4(union Cell1105_4 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL1105_4 : KIND_EMPTY1105_4;
}
struct Node1105_5 {
    int value;
    struct Node1105_5 *next;
};

int length1105_5(const struct Node1105_5 *head) {
    int count = 0;
    const struct Node1105_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
