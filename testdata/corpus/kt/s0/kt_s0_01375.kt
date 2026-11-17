package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter11330 {
    var current: Int by Delegates.observable(87) { _, _, new -> println(new) }

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
object Registry11331 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11331 = "registry-94"
    }
}
sealed class Outcome11332 {
    data class Success(val payload: String) : Outcome11332()
    data object Failure : Outcome11332()
}

fun describe11332(outcome: Outcome11332): String = when (outcome) {
    is Outcome11332.Success -> outcome.payload
    Outcome11332.Failure -> "failed"
}
