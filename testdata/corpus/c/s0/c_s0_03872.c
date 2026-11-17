
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2619_0 {
    int value;
    struct Node2619_0 *next;
};

int length2619_0(const struct Node2619_0 *head) {
    int count = 0;
    const struct Node2619_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
enum Kind2619_1 { KIND_EMPTY2619_1 = 0, KIND_FULL2619_1 = 12 };

union Cell2619_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2619_1(union Cell2619_1 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL2619_1 : KIND_EMPTY2619_1;
}
struct Node2619_2 {
    int value;
    struct Node2619_2 *next;
};

int length2619_2(const struct Node2619_2 *head) {
    int count = 0;
    const struct Node2619_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
enum Kind2619_3 { KIND_EMPTY2619_3 = 0, KIND_FULL2619_3 = 26 };

union Cell2619_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2619_3(union Cell2619_3 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL2619_3 : KIND_EMPTY2619_3;
}
int scan2619_4(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 33;
}
#define MASK2619_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2619_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2619_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
