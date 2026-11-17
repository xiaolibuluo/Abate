
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1348_0 {
    int value;
    struct Node1348_0 *next;
};

int length1348_0(const struct Node1348_0 *head) {
    int count = 0;
    const struct Node1348_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
#define MASK1348_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1348_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1348_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1348_2 { KIND_EMPTY1348_2 = 0, KIND_FULL1348_2 = 83 };

union Cell1348_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1348_2(union Cell1348_2 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL1348_2 : KIND_EMPTY1348_2;
}
enum Kind1348_3 { KIND_EMPTY1348_3 = 0, KIND_FULL1348_3 = 90 };

union Cell1348_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1348_3(union Cell1348_3 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL1348_3 : KIND_EMPTY1348_3;
}
#if defined(__STDC_VERSION__)
int variant1348_4(int value, ...) {
    return value + 97;
}
#else
int variant1348_4(int value) {
    return value + 97;
}
#endif
struct Node1348_5 {
    int value;
    struct Node1348_5 *next;
};

int length1348_5(const struct Node1348_5 *head) {
    int count = 0;
    const struct Node1348_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
