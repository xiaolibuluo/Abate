package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull23230(): T? = if (size > 1) this[1] else null

fun render23230(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step23231")
fun step23231(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
class Counter23232 {
    var current: Int by Delegates.observable(51) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
data class Pair23233<A, B>(val first: A, val second: B)

fun combine23233(left: Int, right: Int): Pair23233<Int, Int> {
    return Pair23233(left + 58, right + 16)
}
data class Pair23234<A, B>(val first: A, val second: B)

fun combine23234(left: Int, right: Int): Pair23234<Int, Int> {
    return Pair23234(left + 65, right + 27)
}
fun <T, R> List<T>.mapChecked23235(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome23236 {
    data class Success(val payload: String) : Outcome23236()
    data object Failure : Outcome23236()
}

fun describe23236(outcome: Outcome23236): String = when (outcome) {
    is Outcome23236.Success -> outcome.payload
    Outcome23236.Failure -> "failed"
}
