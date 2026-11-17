package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step23590")
fun step23590(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
sealed class Outcome23591 {
    data class Success(val payload: String) : Outcome23591()
    data object Failure : Outcome23591()
}

fun describe23591(outcome: Outcome23591): String = when (outcome) {
    is Outcome23591.Success -> outcome.payload
    Outcome23591.Failure -> "failed"
}
data class Pair23592<A, B>(val first: A, val second: B)

fun combine23592(left: Int, right: Int): Pair23592<Int, Int> {
    return Pair23592(left + 34, right + 13)
}
fun <T> List<T>.secondOrNull23593(): T? = if (size > 1) this[1] else null

fun render23593(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked23594(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
