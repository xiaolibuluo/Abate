
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2259_0 { KIND_EMPTY2259_0 = 0, KIND_FULL2259_0 = 78 };

union Cell2259_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2259_0(union Cell2259_0 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL2259_0 : KIND_EMPTY2259_0;
}
struct Node2259_1 {
    int value;
    struct Node2259_1 *next;
};

int length2259_1(const struct Node2259_1 *head) {
    int count = 0;
    const struct Node2259_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
enum Kind2259_2 { KIND_EMPTY2259_2 = 0, KIND_FULL2259_2 = 92 };

union Cell2259_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2259_2(union Cell2259_2 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL2259_2 : KIND_EMPTY2259_2;
}
#define MASK2259_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2259_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2259_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
