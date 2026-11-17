package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8990<A, B>(val first: A, val second: B)

fun combine8990(left: Int, right: Int): Pair8990<Int, Int> {
    return Pair8990(left + 52, right + 19)
}
fun <T, R> List<T>.mapChecked8991(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome8992 {
    data class Success(val payload: String) : Outcome8992()
    data object Failure : Outcome8992()
}

fun describe8992(outcome: Outcome8992): String = when (outcome) {
    is Outcome8992.Success -> outcome.payload
    Outcome8992.Failure -> "failed"
}
sealed class Outcome8993 {
    data class Success(val payload: String) : Outcome8993()
    data object Failure : Outcome8993()
}

fun describe8993(outcome: Outcome8993): String = when (outcome) {
    is Outcome8993.Success -> outcome.payload
    Outcome8993.Failure -> "failed"
}
data class Pair8994<A, B>(val first: A, val second: B)

fun combine8994(left: Int, right: Int): Pair8994<Int, Int> {
    return Pair8994(left + 80, right + 63)
}
