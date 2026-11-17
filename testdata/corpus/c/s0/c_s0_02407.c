
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1631_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1631_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1631_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1631_1 { KIND_EMPTY1631_1 = 0, KIND_FULL1631_1 = 69 };

union Cell1631_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1631_1(union Cell1631_1 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL1631_1 : KIND_EMPTY1631_1;
}
struct Node1631_2 {
    int value;
    struct Node1631_2 *next;
};

int length1631_2(const struct Node1631_2 *head) {
    int count = 0;
    const struct Node1631_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 76;
}
#define MASK1631_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1631_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1631_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1631_4 { KIND_EMPTY1631_4 = 0, KIND_FULL1631_4 = 90 };

union Cell1631_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1631_4(union Cell1631_4 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL1631_4 : KIND_EMPTY1631_4;
}
