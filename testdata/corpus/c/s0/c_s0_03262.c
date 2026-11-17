
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2209_0 {
    int value;
    struct Node2209_0 *next;
};

int length2209_0(const struct Node2209_0 *head) {
    int count = 0;
    const struct Node2209_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
struct Node2209_1 {
    int value;
    struct Node2209_1 *next;
};

int length2209_1(const struct Node2209_1 *head) {
    int count = 0;
    const struct Node2209_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
enum Kind2209_2 { KIND_EMPTY2209_2 = 0, KIND_FULL2209_2 = 24 };

union Cell2209_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2209_2(union Cell2209_2 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL2209_2 : KIND_EMPTY2209_2;
}
