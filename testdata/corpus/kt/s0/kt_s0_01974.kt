package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull16320(): T? = if (size > 1) this[1] else null

fun render16320(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry16321 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16321 = "registry-82"
    }
}
@JvmName("step16322")
fun step16322(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 89
}
sealed class Outcome16323 {
    data class Success(val payload: String) : Outcome16323()
    data object Failure : Outcome16323()
}

fun describe16323(outcome: Outcome16323): String = when (outcome) {
    is Outcome16323.Success -> outcome.payload
    Outcome16323.Failure -> "failed"
}
object Registry16324 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16324 = "registry-6"
    }
}
object Registry16325 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16325 = "registry-13"
    }
}
