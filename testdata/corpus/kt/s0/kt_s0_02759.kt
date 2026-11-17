package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome22780 {
    data class Success(val payload: String) : Outcome22780()
    data object Failure : Outcome22780()
}

fun describe22780(outcome: Outcome22780): String = when (outcome) {
    is Outcome22780.Success -> outcome.payload
    Outcome22780.Failure -> "failed"
}
class Counter22781 {
    var current: Int by Delegates.observable(41) { _, _, new -> println(new) }

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
@JvmName("step22782")
fun step22782(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
sealed class Outcome22783 {
    data class Success(val payload: String) : Outcome22783()
    data object Failure : Outcome22783()
}

fun describe22783(outcome: Outcome22783): String = when (outcome) {
    is Outcome22783.Success -> outcome.payload
    Outcome22783.Failure -> "failed"
}
object Registry22784 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22784 = "registry-62"
    }
}
data class Pair22785<A, B>(val first: A, val second: B)

fun combine22785(left: Int, right: Int): Pair22785<Int, Int> {
    return Pair22785(left + 69, right + 28)
}
class Counter22786 {
    var current: Int by Delegates.observable(76) { _, _, new -> println(new) }

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
