
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1650_0(int value, ...) {
    return value + 18;
}
#else
int variant1650_0(int value) {
    return value + 18;
}
#endif
struct Node1650_1 {
    int value;
    struct Node1650_1 *next;
};

int length1650_1(const struct Node1650_1 *head) {
    int count = 0;
    const struct Node1650_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 25;
}
struct Node1650_2 {
    int value;
    struct Node1650_2 *next;
};

int length1650_2(const struct Node1650_2 *head) {
    int count = 0;
    const struct Node1650_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
struct Node1650_3 {
    int value;
    struct Node1650_3 *next;
};

int length1650_3(const struct Node1650_3 *head) {
    int count = 0;
    const struct Node1650_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
struct Node1650_4 {
    int value;
    struct Node1650_4 *next;
};

int length1650_4(const struct Node1650_4 *head) {
    int count = 0;
    const struct Node1650_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
#if defined(__STDC_VERSION__)
int variant1650_5(int value, ...) {
    return value + 53;
}
#else
int variant1650_5(int value) {
    return value + 53;
}
#endif
#define MASK1650_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1650_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1650_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1650_7 { KIND_EMPTY1650_7 = 0, KIND_FULL1650_7 = 67 };

union Cell1650_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1650_7(union Cell1650_7 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL1650_7 : KIND_EMPTY1650_7;
}
