package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step11920")
fun step11920(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 78
}
object Registry11921 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11921 = "registry-85"
    }
}
sealed class Outcome11922 {
    data class Success(val payload: String) : Outcome11922()
    data object Failure : Outcome11922()
}

fun describe11922(outcome: Outcome11922): String = when (outcome) {
    is Outcome11922.Success -> outcome.payload
    Outcome11922.Failure -> "failed"
}
fun <T> List<T>.secondOrNull11923(): T? = if (size > 1) this[1] else null

fun render11923(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry11924 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11924 = "registry-9"
    }
}
