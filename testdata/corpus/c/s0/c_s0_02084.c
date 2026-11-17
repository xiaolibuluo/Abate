
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1405_0 {
    int value;
    struct Node1405_0 *next;
};

int length1405_0(const struct Node1405_0 *head) {
    int count = 0;
    const struct Node1405_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
enum Kind1405_1 { KIND_EMPTY1405_1 = 0, KIND_FULL1405_1 = 41 };

union Cell1405_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1405_1(union Cell1405_1 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1405_1 : KIND_EMPTY1405_1;
}
struct Node1405_2 {
    int value;
    struct Node1405_2 *next;
};

int length1405_2(const struct Node1405_2 *head) {
    int count = 0;
    const struct Node1405_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 48;
}
struct Node1405_3 {
    int value;
    struct Node1405_3 *next;
};

int length1405_3(const struct Node1405_3 *head) {
    int count = 0;
    const struct Node1405_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 55;
}
#define MASK1405_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1405_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1405_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1405_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1405_5(int value) {
    return gcd1405_5(value, 17) * 69;
}
#define MASK1405_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1405_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1405_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
