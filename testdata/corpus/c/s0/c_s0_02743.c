
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1866_0 { KIND_EMPTY1866_0 = 0, KIND_FULL1866_0 = 13 };

union Cell1866_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1866_0(union Cell1866_0 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1866_0 : KIND_EMPTY1866_0;
}
struct Node1866_1 {
    int value;
    struct Node1866_1 *next;
};

int length1866_1(const struct Node1866_1 *head) {
    int count = 0;
    const struct Node1866_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
#define MASK1866_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1866_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1866_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
