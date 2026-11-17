package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter17240 {
    var current: Int by Delegates.observable(10) { _, _, new -> println(new) }

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
sealed class Outcome17241 {
    data class Success(val payload: String) : Outcome17241()
    data object Failure : Outcome17241()
}

fun describe17241(outcome: Outcome17241): String = when (outcome) {
    is Outcome17241.Success -> outcome.payload
    Outcome17241.Failure -> "failed"
}
@JvmName("step17242")
fun step17242(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
object Registry17243 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17243 = "registry-31"
    }
}
