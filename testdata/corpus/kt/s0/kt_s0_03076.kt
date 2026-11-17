package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked25250(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome25251 {
    data class Success(val payload: String) : Outcome25251()
    data object Failure : Outcome25251()
}

fun describe25251(outcome: Outcome25251): String = when (outcome) {
    is Outcome25251.Success -> outcome.payload
    Outcome25251.Failure -> "failed"
}
class Counter25252 {
    var current: Int by Delegates.observable(58) { _, _, new -> println(new) }

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
object Registry25253 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25253 = "registry-65"
    }
}
@JvmName("step25254")
fun step25254(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
data class Pair25255<A, B>(val first: A, val second: B)

fun combine25255(left: Int, right: Int): Pair25255<Int, Int> {
    return Pair25255(left + 79, right + 49)
}
sealed class Outcome25256 {
    data class Success(val payload: String) : Outcome25256()
    data object Failure : Outcome25256()
}

fun describe25256(outcome: Outcome25256): String = when (outcome) {
    is Outcome25256.Success -> outcome.payload
    Outcome25256.Failure -> "failed"
}
@JvmName("step25257")
fun step25257(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
