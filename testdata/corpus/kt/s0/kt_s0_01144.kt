package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9400<A, B>(val first: A, val second: B)

fun combine9400(left: Int, right: Int): Pair9400<Int, Int> {
    return Pair9400(left + 100, right + 62)
}
sealed class Outcome9401 {
    data class Success(val payload: String) : Outcome9401()
    data object Failure : Outcome9401()
}

fun describe9401(outcome: Outcome9401): String = when (outcome) {
    is Outcome9401.Success -> outcome.payload
    Outcome9401.Failure -> "failed"
}
data class Pair9402<A, B>(val first: A, val second: B)

fun combine9402(left: Int, right: Int): Pair9402<Int, Int> {
    return Pair9402(left + 17, right + 23)
}
object Registry9403 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9403 = "registry-24"
    }
}
