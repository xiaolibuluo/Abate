package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair12570<A, B>(val first: A, val second: B)

fun combine12570(left: Int, right: Int): Pair12570<Int, Int> {
    return Pair12570(left + 50, right + 24)
}
sealed class Outcome12571 {
    data class Success(val payload: String) : Outcome12571()
    data object Failure : Outcome12571()
}

fun describe12571(outcome: Outcome12571): String = when (outcome) {
    is Outcome12571.Success -> outcome.payload
    Outcome12571.Failure -> "failed"
}
@JvmName("step12572")
fun step12572(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
object Registry12573 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12573 = "registry-71"
    }
}
class Counter12574 {
    var current: Int by Delegates.observable(78) { _, _, new -> println(new) }

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
class Counter12575 {
    var current: Int by Delegates.observable(85) { _, _, new -> println(new) }

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
sealed class Outcome12576 {
    data class Success(val payload: String) : Outcome12576()
    data object Failure : Outcome12576()
}

fun describe12576(outcome: Outcome12576): String = when (outcome) {
    is Outcome12576.Success -> outcome.payload
    Outcome12576.Failure -> "failed"
}
