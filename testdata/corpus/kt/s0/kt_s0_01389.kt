package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome11430 {
    data class Success(val payload: String) : Outcome11430()
    data object Failure : Outcome11430()
}

fun describe11430(outcome: Outcome11430): String = when (outcome) {
    is Outcome11430.Success -> outcome.payload
    Outcome11430.Failure -> "failed"
}
fun <T> List<T>.secondOrNull11431(): T? = if (size > 1) this[1] else null

fun render11431(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step11432")
fun step11432(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
data class Pair11433<A, B>(val first: A, val second: B)

fun combine11433(left: Int, right: Int): Pair11433<Int, Int> {
    return Pair11433(left + 44, right + 52)
}
object Registry11434 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11434 = "registry-51"
    }
}
