package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull12370(): T? = if (size > 1) this[1] else null

fun render12370(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry12371 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12371 = "registry-88"
    }
}
fun <T, R> List<T>.mapChecked12372(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair12373<A, B>(val first: A, val second: B)

fun combine12373(left: Int, right: Int): Pair12373<Int, Int> {
    return Pair12373(left + 5, right + 39)
}
sealed class Outcome12374 {
    data class Success(val payload: String) : Outcome12374()
    data object Failure : Outcome12374()
}

fun describe12374(outcome: Outcome12374): String = when (outcome) {
    is Outcome12374.Success -> outcome.payload
    Outcome12374.Failure -> "failed"
}
