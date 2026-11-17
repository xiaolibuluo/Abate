package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16630<A, B>(val first: A, val second: B)

fun combine16630(left: Int, right: Int): Pair16630<Int, Int> {
    return Pair16630(left + 90, right + 60)
}
object Registry16631 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16631 = "registry-97"
    }
}
sealed class Outcome16632 {
    data class Success(val payload: String) : Outcome16632()
    data object Failure : Outcome16632()
}

fun describe16632(outcome: Outcome16632): String = when (outcome) {
    is Outcome16632.Success -> outcome.payload
    Outcome16632.Failure -> "failed"
}
