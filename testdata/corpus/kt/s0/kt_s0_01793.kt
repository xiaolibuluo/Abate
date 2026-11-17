package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry14720 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14720 = "registry-32"
    }
}
sealed class Outcome14721 {
    data class Success(val payload: String) : Outcome14721()
    data object Failure : Outcome14721()
}

fun describe14721(outcome: Outcome14721): String = when (outcome) {
    is Outcome14721.Success -> outcome.payload
    Outcome14721.Failure -> "failed"
}
class Counter14722 {
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
