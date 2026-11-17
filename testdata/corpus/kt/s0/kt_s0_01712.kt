package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14070")
fun step14070(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
sealed class Outcome14071 {
    data class Success(val payload: String) : Outcome14071()
    data object Failure : Outcome14071()
}

fun describe14071(outcome: Outcome14071): String = when (outcome) {
    is Outcome14071.Success -> outcome.payload
    Outcome14071.Failure -> "failed"
}
sealed class Outcome14072 {
    data class Success(val payload: String) : Outcome14072()
    data object Failure : Outcome14072()
}

fun describe14072(outcome: Outcome14072): String = when (outcome) {
    is Outcome14072.Success -> outcome.payload
    Outcome14072.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked14073(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step14074")
fun step14074(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 88
}
@JvmName("step14075")
fun step14075(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
object Registry14076 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14076 = "registry-5"
    }
}
fun <T> List<T>.secondOrNull14077(): T? = if (size > 1) this[1] else null

fun render14077(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull14078(): T? = if (size > 1) this[1] else null

fun render14078(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
