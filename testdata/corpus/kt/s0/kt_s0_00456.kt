package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3730<A, B>(val first: A, val second: B)

fun combine3730(left: Int, right: Int): Pair3730<Int, Int> {
    return Pair3730(left + 101, right + 58)
}
class Counter3731 {
    var current: Int by Delegates.observable(11) { _, _, new -> println(new) }

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
sealed class Outcome3732 {
    data class Success(val payload: String) : Outcome3732()
    data object Failure : Outcome3732()
}

fun describe3732(outcome: Outcome3732): String = when (outcome) {
    is Outcome3732.Success -> outcome.payload
    Outcome3732.Failure -> "failed"
}
@JvmName("step3733")
fun step3733(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 25
}
data class Pair3734<A, B>(val first: A, val second: B)

fun combine3734(left: Int, right: Int): Pair3734<Int, Int> {
    return Pair3734(left + 32, right + 41)
}
object Registry3735 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3735 = "registry-39"
    }
}
class Counter3736 {
    var current: Int by Delegates.observable(46) { _, _, new -> println(new) }

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
