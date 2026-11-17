package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked20860(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair20861<A, B>(val first: A, val second: B)

fun combine20861(left: Int, right: Int): Pair20861<Int, Int> {
    return Pair20861(left + 67, right + 43)
}
@JvmName("step20862")
fun step20862(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
@JvmName("step20863")
fun step20863(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
fun <T, R> List<T>.mapChecked20864(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter20865 {
    var current: Int by Delegates.observable(95) { _, _, new -> println(new) }

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
sealed class Outcome20866 {
    data class Success(val payload: String) : Outcome20866()
    data object Failure : Outcome20866()
}

fun describe20866(outcome: Outcome20866): String = when (outcome) {
    is Outcome20866.Success -> outcome.payload
    Outcome20866.Failure -> "failed"
}
@JvmName("step20867")
fun step20867(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}
fun <T, R> List<T>.mapChecked20868(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
