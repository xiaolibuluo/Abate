package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12510 {
    data class Success(val payload: String) : Outcome12510()
    data object Failure : Outcome12510()
}

fun describe12510(outcome: Outcome12510): String = when (outcome) {
    is Outcome12510.Success -> outcome.payload
    Outcome12510.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked12511(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull12512(): T? = if (size > 1) this[1] else null

fun render12512(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair12513<A, B>(val first: A, val second: B)

fun combine12513(left: Int, right: Int): Pair12513<Int, Int> {
    return Pair12513(left + 90, right + 15)
}
data class Pair12514<A, B>(val first: A, val second: B)

fun combine12514(left: Int, right: Int): Pair12514<Int, Int> {
    return Pair12514(left + 97, right + 26)
}
fun <T> List<T>.secondOrNull12515(): T? = if (size > 1) this[1] else null

fun render12515(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step12516")
fun step12516(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
