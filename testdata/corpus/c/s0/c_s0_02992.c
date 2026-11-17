
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2024_0 { KIND_EMPTY2024_0 = 0, KIND_FULL2024_0 = 30 };

union Cell2024_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2024_0(union Cell2024_0 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL2024_0 : KIND_EMPTY2024_0;
}
#define MASK2024_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2024_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2024_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2024_2 {
    int value;
    struct Node2024_2 *next;
};

int length2024_2(const struct Node2024_2 *head) {
    int count = 0;
    const struct Node2024_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
struct Node2024_3 {
    int value;
    struct Node2024_3 *next;
};

int length2024_3(const struct Node2024_3 *head) {
    int count = 0;
    const struct Node2024_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}
enum Kind2024_4 { KIND_EMPTY2024_4 = 0, KIND_FULL2024_4 = 58 };

union Cell2024_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2024_4(union Cell2024_4 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL2024_4 : KIND_EMPTY2024_4;
}
