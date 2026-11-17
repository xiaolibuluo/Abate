
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2642_0 { KIND_EMPTY2642_0 = 0, KIND_FULL2642_0 = 13 };

union Cell2642_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2642_0(union Cell2642_0 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL2642_0 : KIND_EMPTY2642_0;
}
enum Kind2642_1 { KIND_EMPTY2642_1 = 0, KIND_FULL2642_1 = 20 };

union Cell2642_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2642_1(union Cell2642_1 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL2642_1 : KIND_EMPTY2642_1;
}
struct Node2642_2 {
    int value;
    struct Node2642_2 *next;
};

int length2642_2(const struct Node2642_2 *head) {
    int count = 0;
    const struct Node2642_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
