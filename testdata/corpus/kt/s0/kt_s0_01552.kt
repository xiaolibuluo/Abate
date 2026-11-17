package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12860 {
    data class Success(val payload: String) : Outcome12860()
    data object Failure : Outcome12860()
}

fun describe12860(outcome: Outcome12860): String = when (outcome) {
    is Outcome12860.Success -> outcome.payload
    Outcome12860.Failure -> "failed"
}
data class Pair12861<A, B>(val first: A, val second: B)

fun combine12861(left: Int, right: Int): Pair12861<Int, Int> {
    return Pair12861(left + 46, right + 55)
}
fun <T, R> List<T>.mapChecked12862(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull12863(): T? = if (size > 1) this[1] else null

fun render12863(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step12864")
fun step12864(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 67
}
