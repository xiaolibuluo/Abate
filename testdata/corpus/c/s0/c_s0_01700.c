
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1143_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1143_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1143_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1143_1 {
    int value;
    struct Node1143_1 *next;
};

int length1143_1(const struct Node1143_1 *head) {
    int count = 0;
    const struct Node1143_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
struct Node1143_2 {
    int value;
    struct Node1143_2 *next;
};

int length1143_2(const struct Node1143_2 *head) {
    int count = 0;
    const struct Node1143_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
enum Kind1143_3 { KIND_EMPTY1143_3 = 0, KIND_FULL1143_3 = 44 };

union Cell1143_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1143_3(union Cell1143_3 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL1143_3 : KIND_EMPTY1143_3;
}
