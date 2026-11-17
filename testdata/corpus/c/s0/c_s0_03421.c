
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2312_0 {
    int value;
    struct Node2312_0 *next;
};

int length2312_0(const struct Node2312_0 *head) {
    int count = 0;
    const struct Node2312_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
enum Kind2312_1 { KIND_EMPTY2312_1 = 0, KIND_FULL2312_1 = 95 };

union Cell2312_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2312_1(union Cell2312_1 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL2312_1 : KIND_EMPTY2312_1;
}
enum Kind2312_2 { KIND_EMPTY2312_2 = 0, KIND_FULL2312_2 = 5 };

union Cell2312_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2312_2(union Cell2312_2 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL2312_2 : KIND_EMPTY2312_2;
}
#define MASK2312_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2312_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2312_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
