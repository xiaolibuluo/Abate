package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22980<A, B>(val first: A, val second: B)

fun combine22980(left: Int, right: Int): Pair22980<Int, Int> {
    return Pair22980(left + 100, right + 52)
}
fun <T, R> List<T>.mapChecked22981(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome22982 {
    data class Success(val payload: String) : Outcome22982()
    data object Failure : Outcome22982()
}

fun describe22982(outcome: Outcome22982): String = when (outcome) {
    is Outcome22982.Success -> outcome.payload
    Outcome22982.Failure -> "failed"
}
