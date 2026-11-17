package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome13820 {
    data class Success(val payload: String) : Outcome13820()
    data object Failure : Outcome13820()
}

fun describe13820(outcome: Outcome13820): String = when (outcome) {
    is Outcome13820.Success -> outcome.payload
    Outcome13820.Failure -> "failed"
}
object Registry13821 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13821 = "registry-33"
    }
}
class Counter13822 {
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
sealed class Outcome13823 {
    data class Success(val payload: String) : Outcome13823()
    data object Failure : Outcome13823()
}

fun describe13823(outcome: Outcome13823): String = when (outcome) {
    is Outcome13823.Success -> outcome.payload
    Outcome13823.Failure -> "failed"
}
