package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20630<A, B>(val first: A, val second: B)

fun combine20630(left: Int, right: Int): Pair20630<Int, Int> {
    return Pair20630(left + 52, right + 54)
}
sealed class Outcome20631 {
    data class Success(val payload: String) : Outcome20631()
    data object Failure : Outcome20631()
}

fun describe20631(outcome: Outcome20631): String = when (outcome) {
    is Outcome20631.Success -> outcome.payload
    Outcome20631.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked20632(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry20633 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20633 = "registry-73"
    }
}
