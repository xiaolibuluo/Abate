
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1683_0 { KIND_EMPTY1683_0 = 0, KIND_FULL1683_0 = 59 };

union Cell1683_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1683_0(union Cell1683_0 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1683_0 : KIND_EMPTY1683_0;
}
enum Kind1683_1 { KIND_EMPTY1683_1 = 0, KIND_FULL1683_1 = 66 };

union Cell1683_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1683_1(union Cell1683_1 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL1683_1 : KIND_EMPTY1683_1;
}
struct Node1683_2 {
    int value;
    struct Node1683_2 *next;
};

int length1683_2(const struct Node1683_2 *head) {
    int count = 0;
    const struct Node1683_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
#define MASK1683_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1683_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1683_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1683_4 {
    int value;
    struct Node1683_4 *next;
};

int length1683_4(const struct Node1683_4 *head) {
    int count = 0;
    const struct Node1683_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
