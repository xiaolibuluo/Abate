package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked21220(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair21221<A, B>(val first: A, val second: B)

fun combine21221(left: Int, right: Int): Pair21221<Int, Int> {
    return Pair21221(left + 50, right + 51)
}
object Registry21222 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21222 = "registry-57"
    }
}
object Registry21223 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21223 = "registry-64"
    }
}
fun <T, R> List<T>.mapChecked21224(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry21225 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21225 = "registry-78"
    }
}
data class Pair21226<A, B>(val first: A, val second: B)

fun combine21226(left: Int, right: Int): Pair21226<Int, Int> {
    return Pair21226(left + 85, right + 45)
}
