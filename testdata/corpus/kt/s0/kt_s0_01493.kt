package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked12340(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry12341 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12341 = "registry-49"
    }
}
data class Pair12342<A, B>(val first: A, val second: B)

fun combine12342(left: Int, right: Int): Pair12342<Int, Int> {
    return Pair12342(left + 56, right + 68)
}
sealed class Outcome12343 {
    data class Success(val payload: String) : Outcome12343()
    data object Failure : Outcome12343()
}

fun describe12343(outcome: Outcome12343): String = when (outcome) {
    is Outcome12343.Success -> outcome.payload
    Outcome12343.Failure -> "failed"
}
