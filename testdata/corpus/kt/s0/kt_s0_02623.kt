package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry21650 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21650 = "registry-20"
    }
}
class Counter21651 {
    var current: Int by Delegates.observable(27) { _, _, new -> println(new) }

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
sealed class Outcome21652 {
    data class Success(val payload: String) : Outcome21652()
    data object Failure : Outcome21652()
}

fun describe21652(outcome: Outcome21652): String = when (outcome) {
    is Outcome21652.Success -> outcome.payload
    Outcome21652.Failure -> "failed"
}
sealed class Outcome21653 {
    data class Success(val payload: String) : Outcome21653()
    data object Failure : Outcome21653()
}

fun describe21653(outcome: Outcome21653): String = when (outcome) {
    is Outcome21653.Success -> outcome.payload
    Outcome21653.Failure -> "failed"
}
