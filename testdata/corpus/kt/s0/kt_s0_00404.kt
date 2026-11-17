package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter3310 {
    var current: Int by Delegates.observable(40) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull3311(): T? = if (size > 1) this[1] else null

fun render3311(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step3312")
fun step3312(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 54
}
fun <T, R> List<T>.mapChecked3313(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair3314<A, B>(val first: A, val second: B)

fun combine3314(left: Int, right: Int): Pair3314<Int, Int> {
    return Pair3314(left + 68, right + 52)
}
data class Pair3315<A, B>(val first: A, val second: B)

fun combine3315(left: Int, right: Int): Pair3315<Int, Int> {
    return Pair3315(left + 75, right + 63)
}
@JvmName("step3316")
fun step3316(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
sealed class Outcome3317 {
    data class Success(val payload: String) : Outcome3317()
    data object Failure : Outcome3317()
}

fun describe3317(outcome: Outcome3317): String = when (outcome) {
    is Outcome3317.Success -> outcome.payload
    Outcome3317.Failure -> "failed"
}
