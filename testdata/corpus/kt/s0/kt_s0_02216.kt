package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull18240(): T? = if (size > 1) this[1] else null

fun render18240(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18241<A, B>(val first: A, val second: B)

fun combine18241(left: Int, right: Int): Pair18241<Int, Int> {
    return Pair18241(left + 56, right + 39)
}
@JvmName("step18242")
fun step18242(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
fun <T, R> List<T>.mapChecked18243(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome18244 {
    data class Success(val payload: String) : Outcome18244()
    data object Failure : Outcome18244()
}

fun describe18244(outcome: Outcome18244): String = when (outcome) {
    is Outcome18244.Success -> outcome.payload
    Outcome18244.Failure -> "failed"
}
