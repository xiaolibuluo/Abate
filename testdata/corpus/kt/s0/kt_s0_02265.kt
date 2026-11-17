package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter18610 {
    var current: Int by Delegates.observable(45) { _, _, new -> println(new) }

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
object Registry18611 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18611 = "registry-52"
    }
}
sealed class Outcome18612 {
    data class Success(val payload: String) : Outcome18612()
    data object Failure : Outcome18612()
}

fun describe18612(outcome: Outcome18612): String = when (outcome) {
    is Outcome18612.Success -> outcome.payload
    Outcome18612.Failure -> "failed"
}
