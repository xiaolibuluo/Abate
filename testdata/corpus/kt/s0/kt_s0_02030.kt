package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome16790 {
    data class Success(val payload: String) : Outcome16790()
    data object Failure : Outcome16790()
}

fun describe16790(outcome: Outcome16790): String = when (outcome) {
    is Outcome16790.Success -> outcome.payload
    Outcome16790.Failure -> "failed"
}
@JvmName("step16791")
fun step16791(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
@JvmName("step16792")
fun step16792(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
fun <T> List<T>.secondOrNull16793(): T? = if (size > 1) this[1] else null

fun render16793(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair16794<A, B>(val first: A, val second: B)

fun combine16794(left: Int, right: Int): Pair16794<Int, Int> {
    return Pair16794(left + 35, right + 33)
}
fun <T> List<T>.secondOrNull16795(): T? = if (size > 1) this[1] else null

fun render16795(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
