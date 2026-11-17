package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull9020(): T? = if (size > 1) this[1] else null

fun render9020(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome9021 {
    data class Success(val payload: String) : Outcome9021()
    data object Failure : Outcome9021()
}

fun describe9021(outcome: Outcome9021): String = when (outcome) {
    is Outcome9021.Success -> outcome.payload
    Outcome9021.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked9022(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair9023<A, B>(val first: A, val second: B)

fun combine9023(left: Int, right: Int): Pair9023<Int, Int> {
    return Pair9023(left + 15, right + 12)
}
fun <T, R> List<T>.mapChecked9024(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step9025")
fun step9025(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 29
}
