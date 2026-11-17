package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair6060<A, B>(val first: A, val second: B)

fun combine6060(left: Int, right: Int): Pair6060<Int, Int> {
    return Pair6060(left + 26, right + 42)
}
object Registry6061 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6061 = "registry-33"
    }
}
sealed class Outcome6062 {
    data class Success(val payload: String) : Outcome6062()
    data object Failure : Outcome6062()
}

fun describe6062(outcome: Outcome6062): String = when (outcome) {
    is Outcome6062.Success -> outcome.payload
    Outcome6062.Failure -> "failed"
}
data class Pair6063<A, B>(val first: A, val second: B)

fun combine6063(left: Int, right: Int): Pair6063<Int, Int> {
    return Pair6063(left + 47, right + 14)
}
fun <T> List<T>.secondOrNull6064(): T? = if (size > 1) this[1] else null

fun render6064(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
