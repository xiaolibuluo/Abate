package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter24520 {
    var current: Int by Delegates.observable(65) { _, _, new -> println(new) }

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
object Registry24521 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24521 = "registry-72"
    }
}
sealed class Outcome24522 {
    data class Success(val payload: String) : Outcome24522()
    data object Failure : Outcome24522()
}

fun describe24522(outcome: Outcome24522): String = when (outcome) {
    is Outcome24522.Success -> outcome.payload
    Outcome24522.Failure -> "failed"
}
data class Pair24523<A, B>(val first: A, val second: B)

fun combine24523(left: Int, right: Int): Pair24523<Int, Int> {
    return Pair24523(left + 86, right + 65)
}
