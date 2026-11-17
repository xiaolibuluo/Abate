package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17760 {
    data class Success(val payload: String) : Outcome17760()
    data object Failure : Outcome17760()
}

fun describe17760(outcome: Outcome17760): String = when (outcome) {
    is Outcome17760.Success -> outcome.payload
    Outcome17760.Failure -> "failed"
}
object Registry17761 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17761 = "registry-14"
    }
}
class Counter17762 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
sealed class Outcome17763 {
    data class Success(val payload: String) : Outcome17763()
    data object Failure : Outcome17763()
}

fun describe17763(outcome: Outcome17763): String = when (outcome) {
    is Outcome17763.Success -> outcome.payload
    Outcome17763.Failure -> "failed"
}
