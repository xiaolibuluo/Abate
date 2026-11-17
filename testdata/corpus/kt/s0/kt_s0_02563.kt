package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry21100 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21100 = "registry-81"
    }
}
data class Pair21101<A, B>(val first: A, val second: B)

fun combine21101(left: Int, right: Int): Pair21101<Int, Int> {
    return Pair21101(left + 88, right + 28)
}
sealed class Outcome21102 {
    data class Success(val payload: String) : Outcome21102()
    data object Failure : Outcome21102()
}

fun describe21102(outcome: Outcome21102): String = when (outcome) {
    is Outcome21102.Success -> outcome.payload
    Outcome21102.Failure -> "failed"
}
data class Pair21103<A, B>(val first: A, val second: B)

fun combine21103(left: Int, right: Int): Pair21103<Int, Int> {
    return Pair21103(left + 5, right + 50)
}
