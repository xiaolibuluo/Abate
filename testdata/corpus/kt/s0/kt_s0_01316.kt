package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10850 {
    data class Success(val payload: String) : Outcome10850()
    data object Failure : Outcome10850()
}

fun describe10850(outcome: Outcome10850): String = when (outcome) {
    is Outcome10850.Success -> outcome.payload
    Outcome10850.Failure -> "failed"
}
object Registry10851 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10851 = "registry-52"
    }
}
data class Pair10852<A, B>(val first: A, val second: B)

fun combine10852(left: Int, right: Int): Pair10852<Int, Int> {
    return Pair10852(left + 59, right + 62)
}
data class Pair10853<A, B>(val first: A, val second: B)

fun combine10853(left: Int, right: Int): Pair10853<Int, Int> {
    return Pair10853(left + 66, right + 12)
}
